import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { WavesController } from './waves/waves.controller';
import { WavesModule } from './waves/waves.module';

@Module({
  imports: [WavesModule],
  controllers: [AppController, WavesController],
  providers: [AppService],
})
export class AppModule {}
