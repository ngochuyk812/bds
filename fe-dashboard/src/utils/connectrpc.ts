import { createConnectTransport } from '@connectrpc/connect-web';
import { Client, createClient } from '@connectrpc/connect';
import { useMemo } from 'react';
import type { DescService } from '@bufbuild/protobuf';
import { AuthService } from '../proto/genjs/auth/v1/auth_service_pb';
import { PropertyService } from '../proto/genjs/property/v1/property_service_pb';
import { authInterceptor } from './interceptor';

const AUTH_URL = process.env.REACT_APP_API_AUTH_URL || 'https://api-dev.nnh.io.vn/auth-service';
const PROPERTY_URL = process.env.REACT_APP_API_PROPERTY_URL || 'https://api-dev.nnh.io.vn/property-service';


const transportCommonConfig = {
    interceptors: [authInterceptor],
    // useBinaryFormat: true,
};

export const grpcAuthClient = createClient<typeof AuthService>(AuthService, createConnectTransport({
    baseUrl: AUTH_URL,
    ...transportCommonConfig

}));

export const grpcPropertyClient = createClient<typeof PropertyService>(PropertyService, createConnectTransport({
    baseUrl: PROPERTY_URL,
    ...transportCommonConfig

}));
