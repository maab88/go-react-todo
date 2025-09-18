export const api = {
  async get(path: string) {
    const res = await fetch(`/api${path}`)
    if (!res.ok) throw new Error(await res.text())
    return { data: await res.json() }
  },
  async post(path: string, body: any) {
    const res = await fetch(`/api${path}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body)
    })
    if (!res.ok) throw new Error(await res.text())
    return { data: await res.json() }
  },
  async patch(path: string, body: any) {
    const res = await fetch(`/api${path}`, {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body)
    })
    if (!res.ok) throw new Error(await res.text())
    return { data: await res.json() }
  },
  async delete(path: string) {
    const res = await fetch(`/api${path}`, { method: 'DELETE' })
    if (!res.ok) throw new Error(await res.text())
    return {}
  },
}
