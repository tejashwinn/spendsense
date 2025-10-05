"use client"

import { columns } from "@/components/accounts/columns";
import { DataTable } from "@/components/accounts/data-table";
import { Button } from "@/components/ui/button";
import { getAccounts, ModelsAccountResponse } from "@/lib/api";

const accounts: ModelsAccountResponse[] = (await getAccounts<true>())?.data?.items ?? [];
export default function Home() {

  return (
    <>
      <Button>
        Add Account
      </Button>
      <DataTable columns={columns} data={accounts} />
    </>
  )
}
