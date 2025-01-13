import { useEffectOnce } from './useEffectOnce';
import { useLatest } from './useLatest';

export function useUnmount(callback: () => void) {
  const unmountFn = useLatest(callback);

  useEffectOnce(() => unmountFn.current);
}
