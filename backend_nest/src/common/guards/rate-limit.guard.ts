import {
  CanActivate,
  ExecutionContext,
  Injectable,
  HttpException,
} from '@nestjs/common';
import { RedisService } from '../../infra/redis/redis/redis.service';
import Redis from 'ioredis';

@Injectable()
export class RateLimitGuard implements CanActivate {
  private readonly limit: number;
  private readonly redis: Redis;
  constructor(private readonly redisService: RedisService) {
    this.limit = Number(process.env.RATE_LIMIT_PER_HOUR || 1000);
    this.redis = redisService.getClient();
  }

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const req = context.switchToHttp().getRequest();

    // mirror c.Locals("client")
    const client = req.client;
    if (!client) {
      return true;
    }

    // format: YYYY-MM-DD-HH
    const now = new Date();
    const hourKey = now.toISOString().slice(0, 13).replace('T', '-');

    const key = `ratelimit:${client.client_id}:${hourKey}`;

    try {
      const count = await this.redis.incr(key);

      if (count === 1) {
        await this.redis.expire(key, 60 * 60); // 1 hour
      }

      if (count > this.limit) {
        throw new HttpException("Too Many Requests", 429);
      }
    } catch (err) {
      if (err instanceof HttpException) {
        throw err;
      }
      return true;
    }

    return true;
  }
}
