import moment, { Moment } from "moment";

/**
 * Pretty formats the time since upload
 * @param timestamp milliseconds from unix epoch or Moment instance
 */
export const timeSinceUpload = (timestamp: number | Moment) => {
  if (typeof timestamp === "number") timestamp = moment.unix(timestamp);

  const now = moment.utc();
  const duration = moment.duration(now.diff(timestamp));

  const values = [
    ["years", duration.years()],
    ["months", duration.months()],
    ["days", duration.days()],
    ["hours", duration.hours()],
  ]
    .filter(([_, v]) => v !== 0)
    .slice(0, 2) as [string, number][];

  switch (values.length) {
    case 0:
      return "Just now";
    case 1:
      const v = values[0];
      return `${formString(v[0], v[1])} ago`;
    default:
      const [first, second] = values,
        s1 = formString(first[0], first[1]),
        s2 = formString(second[0], second[1]);
      return `${s1} and ${s2} ago`;
  }

  function formString(unit: string, value: number) {
    if (value === 1) unit = unit.slice(0, -1);
    return `${value} ${unit}`;
  }
};
