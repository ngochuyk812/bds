import { createConnectTransport } from '@connectrpc/connect-web';
import { AuthService } from '../proto/genjs/auth/v1/auth_service_pb.js'
import { Client, createClient } from '@connectrpc/connect';
import { useMemo } from 'react';
import { type DescService } from "@bufbuild/protobuf";

const AUTH_URL = process.env.REACT_APP_API_AUTH_URL || 'https://api-dev.nnh.io.vn';


const transport = createConnectTransport({
    baseUrl: AUTH_URL,
});


export function useClient<T extends DescService>(service: T): Client<T> {
    return useMemo(() => createClient(service, transport), [service]);
}