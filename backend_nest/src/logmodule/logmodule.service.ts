import { Injectable } from '@nestjs/common';
import { RedisService } from '../infra/redis/redis/redis.service';
import Redis from 'ioredis';

@Injectable()
export class LogmoduleService {
  private readonly redis: Redis;

  constructor(redisService: RedisService) {
    this.redis = redisService.getClient();
  }

  async InsertLog(clientId: string): Promise<boolean> {
    const date = new Date().toISOString().slice(0, 10);
    const dailyKey = `usage:daily:${clientId}:${date}`;

    await this.redis.incr(dailyKey);
    await this.redis.zincrby('usage:top:24h', 1, clientId);

    return true;
  }
}
