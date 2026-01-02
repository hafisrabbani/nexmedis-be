import { ConflictException, Injectable } from '@nestjs/common';
import { PrismaService } from '../infra/database/prisma/prisma.service';
import { RegisterClientDto } from './dto/register-client.dto';
import {
  encryptEmail,
  generateApiKey,
  hashApiKey,
} from '../common/security/helper.security';

@Injectable()
export class ClientService {
  constructor(private readonly prisma: PrismaService) {}

  async register(dto: RegisterClientDto) {
    try {
      const exist = await this.prisma.clients.findUnique({
        where: { client_id: dto.client_id },
      });

      console.log('exist', exist);

      if (exist) {
        console.log('Conflict: Client ID already registered');
        throw new ConflictException('Client ID already registered');
      }

      const rawApikey = generateApiKey();
      const hashedApiKey = hashApiKey(rawApikey);
      const emailBuffer = encryptEmail(dto.email);
      await this.prisma.clients.create({
        data: {
          name: dto.name,
          client_id: dto.client_id,
          api_key_hash: hashedApiKey,
          email: new Uint8Array(emailBuffer),
        },
      });

      return {
        api_key: rawApikey,
      };
    } catch (error) {
      console.error('Error registering client:', error);
      throw error;
    }
  }
}
