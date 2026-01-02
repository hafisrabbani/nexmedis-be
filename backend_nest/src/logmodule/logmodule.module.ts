import { Module } from '@nestjs/common';
import { LogmoduleService } from './logmodule.service';
import { LogmoduleController } from './logmodule.controller';

@Module({
  controllers: [LogmoduleController],
  providers: [LogmoduleService],
})
export class LogmoduleModule {}
