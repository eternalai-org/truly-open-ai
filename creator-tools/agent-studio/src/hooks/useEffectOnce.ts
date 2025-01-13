import { EffectCallback, useEffect, useRef } from 'react';

export function useEffectOnce(effect: EffectCallback) {
  const destroyFn = useRef<ReturnType<EffectCallback> | null>(null);
  const effectCalled = useRef(false);
  const effectFn = useRef(effect);

  useEffect(() => {
    if (!effectCalled.current) {
      destroyFn.current = effectFn.current();
      effectCalled.current = true;
    }

    return () => {
      if (destroyFn.current) {
        destroyFn.current();
        destroyFn.current = null;
      }
    };
  }, []);
}
