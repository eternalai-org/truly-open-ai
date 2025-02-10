import { useEffect, useRef, useState } from 'react';

import { useUnmount } from './useUnmount';

export function useThrottle<T extends (...arguments_: Array<unknown>) => void>(
  callback: T,
  delayMs = 500,
  trailing: boolean = false,
): () => void {
  const [now, setNow] = useState(0);
  const callbackRef = useRef(callback);
  const hasPendingCall = useRef(false);
  const timer = useRef<number>(undefined as any);

  useEffect(() => {
    callbackRef.current = callback;
  }, [callback]);

  useEffect(() => {
    if (!now) {
      return;
    }

    if (!timer.current) {
      callbackRef.current();

      const timerCallback = () => {
        if (hasPendingCall.current) {
          hasPendingCall.current = false;

          if (trailing) {
            callbackRef.current();
          }

          timer.current = undefined as any;
        } else {
          timer.current = undefined as any;
        }
      };

      timer.current = window.setTimeout(timerCallback, delayMs);
    } else {
      hasPendingCall.current = true;
    }
  }, [delayMs, now, trailing]);

  useUnmount(() => {
    window.clearTimeout(timer.current);
    timer.current = undefined as any;
  });

  return () => setNow(Date.now());
}
