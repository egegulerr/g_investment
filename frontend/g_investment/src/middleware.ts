import { NextRequest, NextResponse } from "next/server";
import { checkJwtValidity } from "./actions/loginActions";

export async function middleware(req: NextRequest) {
  const path = req.nextUrl.pathname;
  const token = req.cookies.get("jwt");

  if (!token && path !== "/login") {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  const isValid = await checkJwtValidity();

  if (!isValid && path !== "/login")
    return NextResponse.redirect(new URL("/login", req.url));
  if (isValid && path === "/") {
    return NextResponse.redirect(new URL("/home", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/", "/login", "/home"],
};
