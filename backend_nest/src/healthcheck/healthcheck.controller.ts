import { Controller, Get } from '@nestjs/common';
import { HealthcheckService } from './healthcheck.service';

@Controller()
export class HealthcheckController {
  constructor(private readonly healthcheckService: HealthcheckService) {}

  @Get('/health')
  check() {
    return this.healthcheckService.check();
  }
}
