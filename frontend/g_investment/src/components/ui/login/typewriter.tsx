"use client";

import { Typewriter } from "react-simple-typewriter";

export default function TypeWriter() {
  return (
    <span className="text-white text-4xl">
      To Make Money
      <br />
      <Typewriter
        words={["Be a G."]}
        loop={false}
        cursor
        cursorStyle="_"
        typeSpeed={70}
        deleteSpeed={50}
        delaySpeed={2000}
      />
    </span>
  );
}
