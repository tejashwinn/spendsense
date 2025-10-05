"use client"
import { CheckIcon, ChevronsUpDownIcon } from "lucide-react"
import * as React from "react"

import { columns } from "@/components/accounts/columns"
import { DataTable } from "@/components/accounts/data-table"
import { Button } from "@/components/ui/button"
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandList,
} from "@/components/ui/command"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import {
  getAccounts,
  getAccountTypes,
  ModelsAccountResponse,
  ModelsAccountTypeResponse,
  ModelsCreateAccountRequest,
  postAccounts
} from "@/lib/api"
import { cn } from "@/lib/utils"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { z } from "zod"
import { client } from "@/lib/api/client.gen"
import { error } from "console"
const accounts: ModelsAccountResponse[] = (await getAccounts<true>())?.data?.items ?? [];
const accountTypes: ModelsAccountTypeResponse[] = (await getAccountTypes<true>())?.data?.items ?? [];
console.log(accountTypes)

const formSchema = z.object({
  name: z.string(),
  balance: z
    .coerce
    .number<number>(),
  provider: z.string(),
  type: z.string()
})


export default function Home() {
  const accountAddForm = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      balance: 0,
      provider: "",
      type: "",
    }
  })

  const [open, setOpen] = React.useState(false)

  function onSubmit(values: z.infer<typeof formSchema>) {
    const accountAddReq: ModelsCreateAccountRequest = {
      name: values.name,
      balance: values.balance,
      provider: values.provider,
      type_id: Number(values.type)
    }

    const response = (postAccounts<true>({
      body: accountAddReq
    })).then((res) => console.log(res.data))
  }



  return (
    <>
      <div className="flex justify-between w-full">
        <div>
          <h1 className="flex items-center space-x-6">Accounts</h1>
        </div>

        <div className="flex gap-2">
          <Dialog>
            <DialogTrigger asChild>
              <Button variant="outline">Add Account</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
              <DialogHeader>
                <DialogTitle>Add Account</DialogTitle>
                <DialogDescription>
                  You can add your account details here
                </DialogDescription>
              </DialogHeader>
              <Form {...accountAddForm}>
                <form onSubmit={accountAddForm.handleSubmit(onSubmit)} className="space-y-4">
                  <FormField
                    control={accountAddForm.control}
                    name="name"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Name</FormLabel>
                        <FormControl>
                          <Input placeholder="Account Name" {...field} />
                        </FormControl>
                        {/* <FormDescription>
                          This will be the account name used for your reference
                        </FormDescription> */}
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={accountAddForm.control}
                    name="provider"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Provider</FormLabel>
                        <FormControl>
                          <Input placeholder="Provider's Name" {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={accountAddForm.control}
                    name="balance"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Balance</FormLabel>
                        <FormControl>
                          <Input
                            type="number"
                            placeholder="Current Balance" {...field}
                          />

                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={accountAddForm.control}
                    name="type"
                    render={({ field }) => (
                      <FormItem className="flex flex-col">
                        <FormLabel>Type</FormLabel>
                        <FormControl>
                          {/* <Input placeholder="Current Balance" {...field} /> */}
                          <Popover open={open} onOpenChange={setOpen}>
                            <PopoverTrigger asChild>
                              <Button
                                variant="outline"
                                role="combobox"
                                aria-expanded={open}
                                className="justify-between"
                              >
                                {
                                  field.value ?
                                    accountTypes.find((accountType) => accountType.id?.toString() === field.value)?.name
                                    : "Select Account Type..."

                                }
                                <ChevronsUpDownIcon className="ml-2 h-4 w-4 shrink-0 opacity-50" />
                              </Button>
                            </PopoverTrigger>
                            <PopoverContent className="w-auto p-0" align="start">
                              <Command>
                                <CommandInput placeholder="Search Account Type..." />
                                <CommandList>
                                  <CommandEmpty>No resuelts</CommandEmpty>
                                  <CommandGroup>
                                    {accountTypes.map((accountType) => (
                                      <CommandItem
                                        key={accountType.id?.toString()}
                                        value={accountType.id?.toString()}
                                        onSelect={(currentValue) => {
                                          field.value === currentValue ? field.onChange("") : field.onChange(currentValue)
                                          setOpen(false)
                                        }}
                                      >
                                        <CheckIcon
                                          className={cn(
                                            "mr-2 h-4 w-4",
                                            field.value === accountType.id?.toString() ? "opacity-100" : "opacity-0"
                                          )}
                                        />
                                        {accountType.name}
                                      </CommandItem>
                                    ))}
                                  </CommandGroup>
                                </CommandList>
                              </Command>
                            </PopoverContent>
                          </Popover>

                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <DialogFooter>
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button type="submit">Submit</Button>
                  </DialogFooter>
                </form>

              </Form>
            </DialogContent>


          </Dialog>
        </div>
      </div >
      <DataTable columns={columns} data={accounts} />
    </>
  )
}
