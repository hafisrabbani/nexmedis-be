import { Body, Controller, Post, Put, Req, UseGuards } from '@nestjs/common';
import { JwtGuard } from '../common/guards/jwt.guard';
import { IpWhitelistService } from './ip-whitelist.service';
import { RateLimitGuard } from '../common/guards/rate-limit.guard';
import { ApiKeyGuard } from '../common/guards/api-key.guard';

@Controller('api')
export class IpWhitelistController {
  constructor(private readonly service: IpWhitelistService) {}

  @Post('whitelist')
  @UseGuards(JwtGuard, RateLimitGuard, ApiKeyGuard)
  async replaceAll(@Req() req, @Body() body: { ips: string[] }) {
    const user = req.user?.client_id;
    console.log(req.user);
    await this.service.replaceAll(user, body.ips ?? []);
    return;
  }
}
