"use client"

import { ColumnDef } from "@tanstack/react-table"
import { ModelsAccountResponse } from "@/lib/api/types.gen"

export const columns: ColumnDef<ModelsAccountResponse>[] = [
    {
        accessorKey: "name",
        header: "Account Name",
    },
    {
        accessorKey: "provider",
        header: "Provider",
    },
    {
        accessorKey: "type.name",
        header: "Type",
        cell: ({ row }) => row.original?.type?.name,
    },
    {
        accessorKey: "currency.code",
        header: "Currency",
        cell: ({ row }) => `${row.original.currency?.code}`,
    },
    {
        accessorKey: "balance",
        header: "Balance",
        cell: ({ row }) => `${row.original.currency?.symbol}${row.original.balance?.toFixed(row.original?.currency?.decimal_places)}`,
    },
    {
        accessorKey: "created_at",
        header: "Created At",
        cell: ({ row }) => row.original.created_at ? new Date(row.original.created_at)?.toLocaleDateString() : "",
    },
]
