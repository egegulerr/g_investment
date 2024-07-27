import { AlignJustify, Gem, House, Youtube } from "lucide-react";
import { Button } from "../ui/button";
import {
  Sheet,
  SheetContent,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "../ui/sheet";
import Link from "next/link";

const categories = [
  {
    id: 1,
    name: "Investment",
    icon: "gem",
    link: "/home/investment",
  },
  {
    id: 2,
    name: "Youtube",
    icon: "youtube",
    link: "/home/youtube",
  },
  {
    id: 3,
    name: "Assets",
    icon: "assets",
    link: "/homehome/assests",
  },
];

enum Icons {
  gem = "gem",
  youtube = "youtube",
  assets = "assets",
}

function Sidebar() {
  function renderIcon(icon: Icons) {
    switch (icon) {
      case Icons.gem:
        return <Gem />;
      case Icons.youtube:
        return <Youtube />;
      case Icons.assets:
        return <House />;
      default:
        return <Gem />;
    }
  }

  return (
    <Sheet key={"left"}>
      <SheetTrigger>
        <AlignJustify>
          <Button></Button>
        </AlignJustify>
      </SheetTrigger>
      <SheetContent side={"left"} className="w-1/6">
        <SheetHeader>
          <SheetTitle>G.</SheetTitle>
          <div className="margin-top-2 grid gap-10 px-2">
            {categories.map((category) => {
              return (
                <div
                  key={category.id}
                  className="hover:bg-primary/90 hover:text-white p-4 rounded-lg flex items-center gap-3"
                >
                  {renderIcon(category.icon as Icons)}
                  <Link href={category.link}>{category.name}</Link>
                </div>
              );
            })}
          </div>
        </SheetHeader>
      </SheetContent>
    </Sheet>
  );
}

export default Sidebar;
