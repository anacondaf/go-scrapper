import {NestFactory} from '@nestjs/core';
import {AppModule} from './app.module';
import {ConfigService} from '@nestjs/config';
import {VersioningType} from '@nestjs/common';
import {MicroserviceOptions, Transport} from '@nestjs/microservices';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  app.setGlobalPrefix('api');

  app.enableVersioning({
    type: VersioningType.URI,
  });

  const configService = app.get(ConfigService);
  const PORT = configService.get('PORT');
  const RABBITMQ_CONNECTION_STRING = configService.get(
      'RABBITMQ_CONNECTION_STRING',
  );

  await app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.RMQ,
    options: {
      urls: [RABBITMQ_CONNECTION_STRING],
      queue: 'hello',
      queueOptions: {
        durable: false,
      },
    },
  });

  await app.startAllMicroservices()

  // await app.listen(PORT);
}

bootstrap();
