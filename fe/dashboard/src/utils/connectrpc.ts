import { createConnectTransport } from '@connectrpc/connect-web';
import { Client, createClient } from '@connectrpc/connect';
import { useMemo } from 'react';
import type { DescService } from '@bufbuild/protobuf';
import { AuthService } from '../proto/genjs/auth/v1/auth_service_pb';

const AUTH_URL = process.env.REACT_APP_API_AUTH_URL || 'https://api-dev.nnh.io.vn';


const transport = createConnectTransport({
    baseUrl: AUTH_URL,
});

export const grpcAuthClient = createClient<typeof AuthService>(AuthService, transport);
