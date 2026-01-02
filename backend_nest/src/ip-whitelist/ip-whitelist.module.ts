import { Module } from '@nestjs/common';
import { IpWhitelistService } from './ip-whitelist.service';
import { IpWhitelistController } from './ip-whitelist.controller';

@Module({
  controllers: [IpWhitelistController],
  providers: [IpWhitelistService],
})
export class IpWhitelistModule {}
