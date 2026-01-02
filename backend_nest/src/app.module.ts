import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { UsageModule } from './usage/usage.module';
import { DatabaseModule } from './infra/database/database.module';
import { RedisModule } from './infra/redis/redis.module';
import { HealthcheckModule } from './healthcheck/healthcheck.module';
import { ClientModule } from './client/client.module';
import { AuthModule } from './auth/auth.module';
import { LogmoduleModule } from './logmodule/logmodule.module';
import { IpWhitelistModule } from './ip-whitelist/ip-whitelist.module';

@Module({
  imports: [
    UsageModule,
    DatabaseModule,
    RedisModule,
    HealthcheckModule,
    ClientModule,
    AuthModule,
    LogmoduleModule,
    IpWhitelistModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
