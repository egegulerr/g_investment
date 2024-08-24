import { getCompanyLogo } from "@/actions/investmentDataActions";
import { useQuery } from "@tanstack/react-query";
import Image from "next/image";

type CompanyLogoProps = {
  symbol: string;
};

export default function CompanyLogo({ symbol }: CompanyLogoProps) {
  const { data: logo, isFetching } = useQuery({
    enabled: !!symbol,
    queryKey: ["companyLogo", symbol],
    queryFn: () => getCompanyLogo(symbol),
  });

  if (isFetching || !logo) {
    return <></>;
  }

  return (
    <div>
      <Image
        src={`data:image/png;base64,${logo}`}
        alt="Logo"
        height={50}
        width={50}
      />
    </div>
  );
}
