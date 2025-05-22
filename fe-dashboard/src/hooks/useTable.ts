import { useState, useEffect, useCallback } from 'react';

export const useTable = <P extends object, R>(
  fn: () => Promise<R>,
) => {
  const [data, setData] = useState<R | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<any>(null);

  const fetchData = async (fn: () => Promise<R>) => {
    try {
      console.log('fetchDataFn type:', typeof fn);

      setLoading(true);
      const result = await fn();
      setData(result);
      setError(null);
    } catch (err) {
      console.error('Fetch error:', err);
      setError(err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchData(fn);
  }, []);


  return {
    data,
    loading,
    error,
    refetch: () => fetchData(fn),
    fetchData,
  };
};
