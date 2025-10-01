export default function DashboardPage() {
  return (
    <main className="p-6">
      <h1 className="text-3xl font-bold mb-4">Spendsense Dashboard</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <section className="bg-white rounded shadow p-4">
          <h2 className="text-xl font-semibold mb-2">Balances Overview</h2>
          {/* Overview of balances will go here */}
        </section>
        <section className="bg-white rounded shadow p-4">
          <h2 className="text-xl font-semibold mb-2">Recent Expenses</h2>
          {/* Recent expenses will go here */}
        </section>
      </div>
      <div className="mt-8">
        <h2 className="text-xl font-semibold mb-2">Quick Actions</h2>
        {/* Quick actions (add expense, add group, etc.) will go here */}
      </div>
    </main>
  );
}
