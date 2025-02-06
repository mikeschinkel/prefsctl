package macprefs

func ToDomainNames(domains []string) (dns []DomainName) {
	dns = make([]DomainName, len(domains))
	for i, domain := range domains {
		dns[i] = DomainName(domain)
	}
	return dns
}
