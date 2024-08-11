import {
  ArrowDownIcon,
  ArrowUpIcon,
  ArrowTopRightIcon,
  ArrowBottomLeftIcon,
} from "@radix-ui/react-icons";

function renderArrow(score: number) {
  if (score > 0 && score <= 20) {
    return <ArrowTopRightIcon className="text-green-500 w-6 h-6" />;
  } else if (score > 20) {
    return <ArrowUpIcon className="text-green-500  w-6 h-6" />;
  } else if (score < 0 && score >= -20) {
    return <ArrowBottomLeftIcon className="text-red-500  w-6 h-6" />;
  } else if (score < -20) {
    return <ArrowDownIcon className="text-red-500  w-6 h-6" />;
  }
}

export default function Score({ score }: { score: number }) {
  const percentage = score * 100;
  return (
    <div className="flex gap-2 font-medium">
      {renderArrow(percentage)}
      <span>{percentage.toFixed(2)}%</span>
    </div>
  );
}
