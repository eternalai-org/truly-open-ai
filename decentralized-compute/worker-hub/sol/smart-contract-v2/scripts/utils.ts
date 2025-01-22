export function combineDurations(
  submitDuration: number,
  commitDuration: number,
  revealDuration: number,
  unstakeDelayTime: number,
  penaltyDuration: number
): BigInt {
  // Validate input (optional but recommended)
  if (
    !Number.isInteger(submitDuration) ||
    submitDuration < 0 ||
    submitDuration >= 2 ** 40 ||
    !Number.isInteger(commitDuration) ||
    commitDuration < 0 ||
    commitDuration >= 2 ** 40 ||
    !Number.isInteger(revealDuration) ||
    revealDuration < 0 ||
    revealDuration >= 2 ** 40 ||
    !Number.isInteger(unstakeDelayTime) ||
    unstakeDelayTime < 0 ||
    unstakeDelayTime >= 2 ** 40 ||
    !Number.isInteger(penaltyDuration) ||
    penaltyDuration < 0 ||
    penaltyDuration >= 2 ** 40
  ) {
    throw new Error(
      "Invalid duration part(s). Each part must be an integer between 0 and 2**40 - 1."
    );
  }

  // Construct the BigNumber representation
  let duration = BigInt(submitDuration.toString()) << BigInt(160); // Shift submitDuration left by 120 bits
  duration = duration + (BigInt(commitDuration.toString()) << BigInt(120)); // Add commitDuration shifted left by 80 bits
  duration = duration + (BigInt(revealDuration.toString()) << BigInt(80)); // Add revealDuration shifted left by 40 bits
  duration = duration + (BigInt(unstakeDelayTime.toString()) << BigInt(40)); // Add unstakeDelayTime directly
  duration = duration + BigInt(penaltyDuration.toString()); // Add unstakeDelayTime directly

  return duration;
}

// Example usage:
// const submitDuration = 111;
// const commitDuration = 222;
// const revealDuration = 333;
// const unstakeDelayTime = 444;
// const penaltyDuration = 555;

// const combinedDuration = combineDurations(
//   submitDuration,
//   commitDuration,
//   revealDuration,
//   unstakeDelayTime,
//   penaltyDuration
// );
// console.log(combinedDuration.toString()); // Output the result as a hexadecimal string (common for BigNumbers)
