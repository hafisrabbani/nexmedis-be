import { Injectable } from '@nestjs/common';
import { RedisService } from '../infra/redis/redis/redis.service';
import Redis from 'ioredis';

@Injectable()
export class UsageService {
  private readonly redis: Redis;
  constructor(private readonly redisService: RedisService) {
    this.redis = this.redisService.getClient();
  }

  async getDaily(clientId: string, days = 7) {
    const result: { date: string; total_requests: number }[] = [];

    for (let i = 0; i < days; i++) {
      const date = new Date();
      date.setDate(date.getDate() - i);

      const day = date.toISOString().slice(0, 10);
      const key = `usage:daily:${clientId}:${day}`;

      const value = await this.redis.get(key);

      result.push({
        date: day,
        total_requests: Number(value ?? 0),
      });
    }

    return result;
  }


  async getTop(limit = 10) {
    try {
      const result = await this.redis.zrevrange(
        'usage:top:24h',
        0,
        limit - 1,
        'WITHSCORES',
      );

      const data: { client_id: string; total_request: number }[] = [];

      for (let i = 0; i < result.length; i += 2) {
        data.push({
          client_id: result[i],
          total_request: Number(result[i + 1]),
        });
      }

      return data;
    } catch {
      return [];
    }
  }
}
