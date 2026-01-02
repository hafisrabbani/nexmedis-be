import { Body, Controller, Post } from '@nestjs/common';
import { ClientService } from './client.service';
import { RegisterClientDto } from './dto/register-client.dto';

@Controller('api')
export class ClientController {
  constructor(private readonly clientService: ClientService) {}

  @Post('register')
  async register(@Body() dto: RegisterClientDto) {
    return this.clientService.register(dto);
  }
}
