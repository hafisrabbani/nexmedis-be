import { Controller, Post, Headers, UseGuards, Req } from '@nestjs/common';
import { AuthService } from './auth.service';
import { ApiKeyGuard } from '../common/guards/api-key.guard';

@Controller('api')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('token')
  @UseGuards(ApiKeyGuard)
  async issueToken(@Req() req) {
    const client = req.client?.api_key;
    console.log('Client from request:', client);
    const token = await this.authService.issueToken(client);
    return token;
  }
}
