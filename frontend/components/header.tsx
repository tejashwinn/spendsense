"use client"

import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuPortal,
    DropdownMenuSeparator,
    DropdownMenuSub,
    DropdownMenuSubContent,
    DropdownMenuSubTrigger,
    DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { NavigationMenu, NavigationMenuItem } from "@radix-ui/react-navigation-menu";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";
import Link from "next/link";
import { NavigationMenuList } from "./ui/navigation-menu";



export function Header() {
    const { theme, setTheme } = useTheme()

    return (
        <NavigationMenu className="p-3 pl-6 pr-6 border-b-1">
            <NavigationMenuList className="flex justify-between w-full">
                <div className="flex"> {/* Group your left-aligned items */}
                    <NavigationMenuItem>
                        <Link href="/" className="flex items-center space-x-2"> {/* Using flexbox for alignment */}
                            {/* // TODO Add it back with apt logo */}
                            {/* <Image
                                src={Logo} // Path to your image in the public directory
                                alt="Home Icon"
                                width={24} // Set appropriate width
                                height={24} // Set appropriate height
                            /> */}
                            <span>SpendSense</span>
                        </Link>
                    </NavigationMenuItem>
                </div>

                <div className="flex gap-2">
                    <NavigationMenuItem>
                        <DropdownMenu>
                            <DropdownMenuTrigger>
                                <Avatar>
                                    <AvatarFallback >T</AvatarFallback>
                                </Avatar>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent className="width-100">
                                <DropdownMenuLabel>My Account</DropdownMenuLabel>
                                <DropdownMenuSeparator />
                                <DropdownMenuSub>
                                    <DropdownMenuSubTrigger>
                                        Theme
                                        {/* TODO Not able to handle the system case, needs revision on how theme is being set */}
                                        {/* {
                                            theme === 'light' ?  
                                            <Sun className="h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90" />
                                            : <Moon className="h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0" />
                                        } */}
                                        <Sun className="h-[1.2rem] w-[1.2rem] scale-100 rotate-0 transition-all dark:scale-0 dark:-rotate-90" />
                                        <Moon className="h-[1.2rem] w-[1.2rem] scale-0 rotate-90 transition-all dark:scale-100 dark:rotate-0" />

                                    </DropdownMenuSubTrigger>
                                    <DropdownMenuPortal>
                                        <DropdownMenuSubContent>
                                            <DropdownMenuItem onClick={() => setTheme("light")}>
                                                Light
                                            </DropdownMenuItem>
                                            <DropdownMenuItem onClick={() => setTheme("dark")}>
                                                Dark
                                            </DropdownMenuItem>
                                            <DropdownMenuItem onClick={() => setTheme("system")}>
                                                System
                                            </DropdownMenuItem>
                                        </DropdownMenuSubContent>
                                    </DropdownMenuPortal>
                                </DropdownMenuSub>
                            </DropdownMenuContent>
                        </DropdownMenu>
                    </NavigationMenuItem>

                </div>

            </NavigationMenuList>

        </NavigationMenu>
    )
}