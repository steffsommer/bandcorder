const MILLIS_PER_SECOND = 1000;
const DIRECTION_INCREASE = 1;
const DIRECTION_REDUCE = -1;

/**
 * Smoothly interpolates a value towards a target at a constant speed.
 *
 * The interpolator moves the current value towards the target value at the specified
 * speed (units per second).
 */
export class LinearInterpolator {
  private speed: number;
  private value = 0;
  private target = 0;
  private targetUpdateMillis = 0;

  constructor(interpolationSpeed: number) {
    if (!interpolationSpeed || interpolationSpeed <= 0) {
      throw new Error(`Speed must be a positive integer, handed was ${interpolationSpeed}`);
    }
    this.speed = interpolationSpeed;
  }

  /**
   * Get an interpolated value, based on the target value, elapsed time and speed.
   */
  getValue(): number {
    if (this.value === this.target) {
      return this.value;
    }
    const secondsPassed = (Date.now() - this.targetUpdateMillis) / MILLIS_PER_SECOND;
    const direction = this.target > this.value ? DIRECTION_INCREASE : DIRECTION_REDUCE;
    this.value = this.value + this.speed * secondsPassed * direction;
    if (
      (direction === DIRECTION_INCREASE && this.value > this.target) ||
      (direction === DIRECTION_REDUCE && this.value < this.target)
    ) {
      this.value = this.target;
    }
    return this.value;
  }

  /**
   * Updates the target value
   */
  setTarget(value: number): void {
    this.target = value;
    this.targetUpdateMillis = Date.now();
  }
}
