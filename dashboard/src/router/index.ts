import React, { FC } from 'react';
import LoginPage from '../pages/auth/Login';
import MetricsPage from '../pages/dashboard/Metrics';

interface RouteObject {
    path: string,
    element: FC<{}>;
    type?: 'public' | 'private';
}

const routePrivates: RouteObject[] = [
    {
        path: '/',
        element: MetricsPage,
    },
];


const routePublics: RouteObject[] = [
    {
        path: '/login',
        element: LoginPage,
    },
];

export default [
    ...routePublics.map(tmp => ({ ...tmp, type: "public" })),
    ...routePrivates.map(tmp => ({ ...tmp, type: "private" }))
];
