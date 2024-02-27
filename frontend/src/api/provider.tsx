import { ReactNode } from 'react';

import {
  MutationCache,
  QueryCache,
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';

export interface AppProvidersProps {
  children: ReactNode;
}

export function AppProviders({ children }: AppProvidersProps) {
    const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                retry: 2,
                refetchOnWindowFocus: false,
            },
        },
        queryCache: new QueryCache({
            onError: (error: any, query: any) => {
                console.log(query, error)
            },
        }),
        mutationCache: new MutationCache({
            onError: (error: any) => {
                console.log(error)
            },
        }),
    });

    return (
        <QueryClientProvider client={queryClient}>
            {children}
        </QueryClientProvider>
    )
}