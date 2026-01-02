import { Injectable } from '@nestjs/common';

@Injectable()
export class HealthcheckService {
  check() {
    return {
      service: 'Nexmedis API',
      status: 'ok',
      uptime: process.uptime(),
      timestamp: new Date().toISOString(),
    };
  }
}
