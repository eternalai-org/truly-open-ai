"use client";
import { useRef } from "react";
import { gsap } from "gsap";
import { useGSAP } from "@gsap/react";

interface IProp {
  counter: number;
  delay?: number;
}

export default function NumberCounter({ counter, delay = 0 }: IProp) {
  const refContent = useRef<HTMLDivElement>(null);
  const refFake = useRef<{ value: number }>({ value: 0 });

  useGSAP(
    () => {
      gsap.to(refFake.current, {
        scrollTrigger: {
          trigger: refContent.current,
          start: "bottom bottom",
        },
        value: counter,
        duration: 2,
        ease: "power3.inOut",
        delay,
        onUpdate: () => {
          if (refContent.current)
            refContent.current.textContent = Math.round(
              refFake.current.value
            ).toString();
        },
      });
    },
    { dependencies: [counter, delay] }
  );
  return <div ref={refContent}>0</div>;
}
