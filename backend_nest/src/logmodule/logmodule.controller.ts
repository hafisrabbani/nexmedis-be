import { Controller, Post, Req, UseGuards } from '@nestjs/common';
import { ApiKeyGuard } from '../common/guards/api-key.guard';
import { LogmoduleService } from './logmodule.service';

@Controller('api')
export class LogmoduleController {
  constructor(private readonly service: LogmoduleService) {}

  @Post('logs')
  @UseGuards(ApiKeyGuard)
  async Log(@Req() req) {
    const client = req.client?.client_id;
    console.log(req.client);
    await this.service.InsertLog(client);
    return;
  }
}
