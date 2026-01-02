import {
  CanActivate,
  ExecutionContext,
  Injectable,
  UnauthorizedException,
} from '@nestjs/common';
import { PrismaService } from '../../infra/database/prisma/prisma.service';
import { hashApiKey } from '../security/helper.security';

@Injectable()
export class ApiKeyGuard implements CanActivate {
  constructor(private readonly prisma: PrismaService) {}

  async canActivate(context: ExecutionContext): Promise<boolean> {
    const req = context.switchToHttp().getRequest();
    const apiKey = req.headers['x-api-key'];
    if (!apiKey) {
      throw new UnauthorizedException("Missing or invalid 'X-API-Key' header");
    }
    const hashedApiKey = hashApiKey(apiKey);
    const client = await this.prisma.clients.findFirst({
      where: {
        api_key_hash: hashedApiKey,
      },
    });

    if (!client) {
      throw new UnauthorizedException('Invalid API Key');
    }
    req.client = client;
    req.client.api_key = apiKey;
    return true;
  }
}
