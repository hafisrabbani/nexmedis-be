import { Controller, Get, Req, Sse, UseGuards } from '@nestjs/common';
import { UsageService } from './usage.service';
import { JwtGuard } from '../common/guards/jwt.guard';
import { ApiKeyGuard } from '../common/guards/api-key.guard';
import { RateLimitGuard } from '../common/guards/rate-limit.guard';
import { catchError, from, interval, switchMap } from 'rxjs';
import { map } from 'rxjs/operators';

@Controller('api')
export class UsageController {
  constructor(private readonly usageService: UsageService) {}

  @Get('usage/daily')
  @UseGuards(JwtGuard, ApiKeyGuard, RateLimitGuard)
  async daily(@Req() req) {
    const client = req.client?.client_id;
    const count = await this.usageService.getDaily(client);
    return {
      daily_usage: count,
    };
  }

  @Get('usage/top')
  @UseGuards(JwtGuard, ApiKeyGuard, RateLimitGuard)
  async top() {
    const topUsage = await this.usageService.getTop();
    return {
      top_usage: topUsage,
    };
  }

  @Sse('usage/stream')
  @UseGuards(JwtGuard, ApiKeyGuard, RateLimitGuard)
  stream(@Req() req) {
    const clientId = req.client?.client_id;
    return interval(1000).pipe(
      switchMap(() =>
        from(this.usageService.getDaily(clientId)).pipe(
          map((data) => ({
            data: {
              success: true,
              data,
            },
          })),
          catchError(() => [
            {
              data: {
                success: false,
                error: 'failed to fetch usage',
              },
            },
          ]),
        ),
      ),
    );
  }
}
