import { Injectable, NotFoundException } from '@nestjs/common';
import { PrismaService } from '../infra/database/prisma/prisma.service';

@Injectable()
export class IpWhitelistService {
  constructor(private readonly prisma: PrismaService) {}

  async replaceAll(clientId: string, ips: string[]): Promise<void> {
    try {
      const clientUUID = await this.findClientIdByClientId(clientId);
      console.log('clientUUID', clientUUID);
      if (!clientUUID) {
        throw new NotFoundException('Not found');
      }
      await this.prisma.$transaction(async (tx) => {
        await tx.client_ip_whitelists.deleteMany({
          where: {
            client_id: clientUUID,
          },
        });

        for (const ip of ips) {
          if (!ip || ip.trim() === '') {
            continue;
          }

          await tx.client_ip_whitelists.create({
            data: {
              client_id: clientUUID,
              ip_address: ip,
            },
          });
        }
      });
    } catch (error) {
      console.error('Error replacing IP whitelist:', error);
      throw error;
    }
  }

  async findClientIdByClientId(clientId: string): Promise<string | null> {
    const client = await this.prisma.clients.findUnique({
      where: { client_id: clientId },
      select: { id: true },
    });

    return client?.id ?? null;
  }
}
