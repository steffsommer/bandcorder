export class TimeUtils {
  /**
   * Calculate the difference between an iso timestamp and now.
   * Returns the difference to the current date in the format mm:ss
   */
  static sinceStr(dateString: string): string {
    const targetDate = new Date(dateString);
    const currentDate = new Date();

    const diffMs = Math.abs(targetDate.getTime() - currentDate.getTime());
    const totalSeconds = Math.floor(diffMs / 1000);

    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;

    const formattedTime = `${minutes.toString().padStart(2, "0")}:${seconds.toString().padStart(2, "0")}`;

    return formattedTime;
  }

  /**
   * Convert seconds into a string with format "mm:ss", not padded
   */
  static toMinutesSecondsStr(seconds: number): string {
    if (seconds <= 0) {
      return "0";
    }
    const minutes = Math.floor(seconds / 60);
    const leftOverSeconds = seconds % 60;
    if (minutes > 0) {
      return `${minutes}:${leftOverSeconds}`;
    }
    return leftOverSeconds.toString();
  }
}
