import { Global, Module } from '@nestjs/common';
import { JwtModule } from '@nestjs/jwt';
import { AuthController } from './auth.controller';
import { AuthService } from './auth.service';

@Global()
@Module({
  imports: [
    JwtModule.registerAsync({
      useFactory: () => {
        const expiresInSeconds = Number(process.env.JWT_EXPIRED_MINUTES ?? 60);
        return {
          secret: process.env.JWT_SECRET || 'defaultSecretKey',
          signOptions: {
            expiresIn: expiresInSeconds * 60, // convert minutes to seconds
          },
        };
      },
    }),
  ],
  controllers: [AuthController],
  providers: [AuthService],
  exports: [JwtModule]
})
export class AuthModule {}
