import { useEffect, useRef, useState } from "react";
import { useUnmount } from "./useUnmount";

export function useThrottleValue<T>(value: T, delayMs: number): T {
  const [throttledValue, setThrottledValue] = useState<T>(value);
  const hasNextValue = useRef(false);
  const nextValue = useRef<T>(null as any);
  const timer = useRef<number>(undefined as any);

  useEffect(() => {
    if (!timer.current) {
      setThrottledValue(value);

      const timeoutCallback = () => {
        if (hasNextValue.current) {
          hasNextValue.current = false;
          setThrottledValue(nextValue.current as T);
          timer.current = window.setTimeout(timeoutCallback, delayMs);
        } else {
          timer.current = undefined as any;
        }
      };

      timer.current = window.setTimeout(timeoutCallback, delayMs);
    } else {
      hasNextValue.current = true;
      nextValue.current = value as any;
    }
  }, [delayMs, value]);

  useUnmount(() => {
    window.clearTimeout(timer.current);
    timer.current = undefined as any;
  });

  return throttledValue;
}
