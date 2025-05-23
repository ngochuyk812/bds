import React, { FC } from 'react';
import LoginPage from '../pages/auth/Login';
import SignUpPage from '../pages/auth/SignUp';
import SitesPage from '../pages/dashboard/Sites';
import MetricsPage from '../pages/dashboard/Metrics';
import AmenitiesPage from '../pages/dashboard/Amenities';

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
    {
        path: '/sites',
        element: SitesPage,
    },
    {
        path: '/amenities',
        element: AmenitiesPage,
    },
];


const routePublics: RouteObject[] = [
    {
        path: '/login',
        element: LoginPage,
    },
    {
        path: '/register',
        element: SignUpPage,
    },
];

export default [
    ...routePublics.map(tmp => ({ ...tmp, type: "public" })),
    ...routePrivates.map(tmp => ({ ...tmp, type: "private" }))
];
