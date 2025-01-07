import { useQueryClient, useQuery } from '@tanstack/vue-query';

export function useWallpapers() {
  const queryClient = useQueryClient();

  const queryWallpapers = useQuery({
    queryKey: ['wallpapers'],
    queryFn: async () => {},
  });

  return {
    queryWallpapers,
  };
}
