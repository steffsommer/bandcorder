export class TimeUtils {
  /**
   * Calculate the difference between an iso timestamp and now.
   * Returns the difference to the current date in the format mm:ss
   */
  static SinceStr(dateString: string): string {
    const targetDate = new Date(dateString);
    const currentDate = new Date();

    const diffMs = Math.abs(targetDate.getTime() - currentDate.getTime());
    const totalSeconds = Math.floor(diffMs / 1000);

    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    const formattedTime = `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;

    return formattedTime;
  }
}
