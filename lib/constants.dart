import 'package:ctlpkg/client/v1/pacman.pbgrpc.dart';
import 'package:flutter/material.dart';
import 'package:grpc/grpc.dart';

const primaryColor = Color(0xFF4079FA);
const secondaryColor = Color(0xFF2A2D3E);
const bgColor = Color(0xFF212332);

const defaultPadding = 16.0;

var chan = ClientChannel(
  'localhost',
  port: 8080,
  options: const ChannelOptions(
    credentials: ChannelCredentials.insecure(),
  ),
);
var stub = PacmanServiceClient(chan);
