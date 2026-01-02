import { Injectable, NotFoundException } from '@nestjs/common';
import { PrismaService } from '../infra/database/prisma/prisma.service';
import { JwtService } from '@nestjs/jwt';
import { hashApiKey } from '../common/security/helper.security';

@Injectable()
export class AuthService {
  constructor(
    private readonly prisma: PrismaService,
    private readonly JwtService: JwtService,
  ) {}

  async issueToken(rawApiKey: string) {
    try {
      const hashedApiKey = hashApiKey(rawApiKey);
      const client = await this.prisma.clients.findFirst({
        where: {
          api_key_hash: hashedApiKey,
        },
        select: {
          client_id: true,
        },
      });

      if (!client) {
        throw new NotFoundException('Client not found');
      }
      const token = this.JwtService.sign({
        client_id: client.client_id,
      });
      return {
        token,
      };
    } catch (error) {
      console.error('Error issuing token:', error);
      throw error;
    }
  }
}
