import Redis from 'ioredis';
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class RedisService {
  private readonly client: Redis;

  constructor(config: ConfigService) {
    this.client = new Redis({
      host: config.get('REDIS_HOST'),
      port: Number(config.get('REDIS_PORT')),
      password: config.get('REDIS_PASSWORD'),
    });
  }

  getClient(): Redis {
    return this.client;
  }
}
