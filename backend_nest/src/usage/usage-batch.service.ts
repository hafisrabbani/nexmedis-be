import { Injectable, OnModuleInit, OnModuleDestroy } from '@nestjs/common';
import { PrismaService } from '../infra/database/prisma/prisma.service';
import { RedisService } from '../infra/redis/redis/redis.service';

@Injectable()
export class UsageBatchService implements OnModuleInit, OnModuleDestroy {
  private timer: NodeJS.Timeout | null = null;

  constructor(
    private readonly prisma: PrismaService,
    private readonly redisService: RedisService,
  ) {}

  onModuleInit() {
    const intervalSeconds = Number(process.env.INTERVAL_INSERT_DATA ?? 30);

    this.timer = setInterval(() => {
      console.info('[UsageBatchService] Flushing usage data to database');
      this.flush().catch(() => {});
    }, intervalSeconds * 1000);
  }

  onModuleDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  }

  private async flush(): Promise<void> {
    const redis = this.redisService.getClient();

    let keys: string[];
    try {
      keys = await redis.keys('usage:daily:*');
    } catch {
      return;
    }

    for (const key of keys) {
      try {
        // key: usage:daily:{clientID}:{YYYY-MM-DD}
        const parts = key.split(':');
        if (parts.length !== 4) continue;

        const clientId = parts[2];
        const date = new Date(parts[3]);
        if (isNaN(date.getTime())) continue;

        const count = Number(await redis.get(key));
        if (Number.isNaN(count)) continue;

        // get client UUID (DB)
        const client = await this.prisma.clients.findUnique({
          where: { client_id: clientId },
          select: { id: true },
        });
        if (!client) continue;

        // UPSERT daily_usage
        await this.prisma.daily_usage.upsert({
          where: {
            client_id_date: {
              client_id: client.id,
              date,
            },
          },
          create: {
            client_id: client.id,
            date,
            total_requests: count,
          },
          update: {
            total_requests: count,
            updated_at: new Date(),
          },
        });
      } catch {
        continue;
      }
    }
  }
}
