import { Module } from '@nestjs/common';
import { UsageService } from './usage.service';
import { UsageController } from './usage.controller';
import { UsageBatchService } from './usage-batch.service';

@Module({
  controllers: [UsageController],
  providers: [UsageService, UsageBatchService],
})
export class UsageModule {}
