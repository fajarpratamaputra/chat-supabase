package supabase

import (
	supa "github.com/nedpals/supabase-go"
)

type Client struct {
	*supa.Client
}

func NewClient() *supa.Client {
	supabaseUrl := "https://yatnvkzinzeofeocuimb.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InlhdG52a3ppbnplb2Zlb2N1aW1iIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDAwMjE3NjgsImV4cCI6MjAxNTU5Nzc2OH0.0gYu_ccSWBXqzbyp6T3p_7uh9z64O3KAQn_tUX66j4c"

	// Create the connection to Supabase
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return supabase
}
