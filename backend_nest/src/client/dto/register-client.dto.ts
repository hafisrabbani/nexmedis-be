import { IsEmail, IsNotEmpty, IsString } from 'class-validator';

export class RegisterClientDto{
  @IsString()
  @IsNotEmpty()
  name: string;

  @IsString()
  @IsNotEmpty()
  client_id: string;

  @IsString()
  @IsNotEmpty()
  @IsEmail()
  email: string;
}
