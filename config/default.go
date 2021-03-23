package config

// DefaultConfig is the default gitleaks configuration. If --config={path-to-config} is set than the config located
// at {path-to-config} will be used. Alternatively, if --repo-config is set then gitleaks will attempt to
// use the config set in a gitleaks.toml or .gitleaks.toml file in the repo that is run with --repo-config set.
const DefaultConfig = `
title = "gitleaks config"

[[rules]]
	description = "AWS Access Key"
	regex = '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
	tags = ["key", "AWS"]

[[rules]]
	description = "AWS Secret Key"
	regex = '''(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]'''
	tags = ["key", "AWS"]

[[rules]]
	description = "AWS MWS key"
	regex = '''amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'''
	tags = ["key", "AWS", "MWS"]

[[rules]]
	description = "Facebook Secret Key"
	regex = '''(?i)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}['\"]'''
	tags = ["key", "Facebook"]

[[rules]]
	description = "Facebook Client ID"
	regex = '''(?i)(facebook|fb)(.{0,20})?['\"][0-9]{13,17}['\"]'''
	tags = ["key", "Facebook"]

[[rules]]
	description = "Twitter Secret Key"
	regex = '''(?i)twitter(.{0,20})?[0-9a-z]{35,44}'''
	tags = ["key", "Twitter"]

[[rules]]
	description = "Twitter Client ID"
	regex = '''(?i)twitter(.{0,20})?[0-9a-z]{18,25}'''
	tags = ["client", "Twitter"]

[[rules]]
	description = "Github"
	regex = '''(?i)github.{0,3}((?i)token|api|key).{0,10}?(?-i)([0-9a-zA-Z]{35,40})'''
	tags = ["key", "Github"]

[[rules]]
	description = "LinkedIn Client ID"
	regex = '''(?i)linkedin(.{0,20})?(?-i)[0-9a-z]{12}'''
	tags = ["client", "LinkedIn"]

[[rules]]
	description = "LinkedIn Secret Key"
	regex = '''(?i)linkedin(.{0,20})?[0-9a-z]{16}'''
	tags = ["secret", "LinkedIn"]

[[rules]]
	description = "Slack"
	regex = '''xox[baprs]-([0-9a-zA-Z]{10,48})?'''
	tags = ["key", "Slack"]

[[rules]]
	description = "Asymmetric Private Key"
	regex = '''-----BEGIN ((EC|PGP|DSA|RSA|OPENSSH) )?PRIVATE KEY( BLOCK)?-----'''
	tags = ["key", "AsymmetricPrivateKey"]

[[rules]]
	description = "Google API key"
	regex = '''AIza[0-9A-Za-z\\-_]{35}'''
	tags = ["key", "Google"]

[[rules]]
	description = "Google (GCP) Service Account"
	regex = '''"type": "service_account"'''
	tags = ["key", "Google"]

[[rules]]
	description = "Heroku API key"
	regex = '''(?i)heroku(.{0,20})?[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'''
	tags = ["key", "Heroku"]

[[rules]]
	description = "MailChimp API key"
	regex = '''(?i)(mailchimp|mc)(.{0,20})?[0-9a-f]{32}-us[0-9]{1,2}'''
	tags = ["key", "Mailchimp"]

[[rules]]
	description = "Mailgun API key"
	regex = '''((?i)(mailgun|mg)(.{0,20})?)?key-[0-9a-z]{32}'''
	tags = ["key", "Mailgun"]

[[rules]]
	description = "PayPal Braintree access token"
	regex = '''access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}'''
	tags = ["key", "Paypal"]

[[rules]]
	description = "Picatic API key"
	regex = '''sk_live_[0-9a-z]{32}'''
	tags = ["key", "Picatic"]

[[rules]]
	description = "SendGrid API Key"
	regex = '''SG\.[\w_]{16,32}\.[\w_]{16,64}'''
	tags = ["key", "SendGrid"]

[[rules]]
	description = "Slack Webhook"
	regex = '''https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8,12}/[a-zA-Z0-9_]{24}'''
	tags = ["key", "slack"]

[[rules]]
	description = "Stripe API key"
	regex = '''(?i)stripe(.{0,20})?[sr]k_live_[0-9a-zA-Z]{24}'''
	tags = ["key", "Stripe"]

[[rules]]
	description = "Square access token"
	regex = '''sq0atp-[0-9A-Za-z\-_]{22}'''
	tags = ["key", "square"]

[[rules]]
	description = "Square OAuth secret"
	regex = '''sq0csp-[0-9A-Za-z\\-_]{43}'''
	tags = ["key", "square"]

[[rules]]
	description = "Twilio API key"
	regex = '''(?i)twilio(.{0,20})?SK[0-9a-f]{32}'''
	tags = ["key", "twilio"]

[[rules]]
	description = "Dynatrace ttoken"
	regex = '''dt0[a-zA-Z]{1}[0-9]{2}\.[A-Z0-9]{24}\.[A-Z0-9]{64}'''
	tags = ["key", "Dynatrace"]

[[rules]]
	description = "Shopify shared secret"
	regex = '''shpss_[a-fA-F0-9]{32}'''
	tags = ["key", "Shopify"]

[[rules]]
	description = "Shopify access token"
	regex = '''shpat_[a-fA-F0-9]{32}'''
	tags = ["key", "Shopify"]

[[rules]]
	description = "Shopify custom app access token"
	regex = '''shpca_[a-fA-F0-9]{32}'''
	tags = ["key", "Shopify"]

[[rules]]
	description = "Shopify private app access token"
	regex = '''shppa_[a-fA-F0-9]{32}'''
	tags = ["key", "Shopify"]

[allowlist]
	description = "Allowlisted files"
	files = ['''^\.?gitleaks.toml$''',
	'''(.*?)(png|jpg|gif|doc|docx|pdf|bin|xls|pyc|zip)$''',
	'''(go.mod|go.sum)$''']

[[rules]]
	description = "AWS Access Key"
	regex = '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
	tags = ["key", "AWS"]

[[rules]]
	description = "AWS cred file info"
	regex = '''(?i)(aws_access_key_id|aws_secret_access_key)(.{0,20})?=.[0-9a-zA-Z\/+]{20,40}'''
	tags = ["AWS"]

[[rules]]
	description = "AWS Secret Key"
	regex = '''(?i)aws(.{0,20})?(?-i)['\"][0-9a-zA-Z\/+]{40}['\"]'''
	tags = ["key", "AWS"]

[[rules]]
	description = "AWS MWS key"
	regex = '''amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}'''
	tags = ["key", "AWS", "MWS"]

[[rules]]
	description = "Facebook Secret Key"
	regex = '''(?i)(facebook|fb)(.{0,20})?(?-i)['\"][0-9a-f]{32}['\"]'''
	tags = ["key", "Facebook"]

[[rules]]
	description = "Facebook Client ID"
	regex = '''(?i)(facebook|fb)(.{0,20})?['\"][0-9]{13,17}['\"]'''
	tags = ["key", "Facebook"]

[[rules]]
	description = "Twitter Secret Key"
	regex = '''(?i)twitter(.{0,20})?['\"][0-9a-z]{35,44}['\"]'''
	tags = ["key", "Twitter"]

[[rules]]
	description = "Twitter Client ID"
	regex = '''(?i)twitter(.{0,20})?['\"][0-9a-z]{18,25}['\"]'''
	tags = ["client", "Twitter"]

[[rules]]
	description = "Github"
	regex = '''(?i)github(.{0,20})?(?-i)['\"][0-9a-zA-Z]{35,40}['\"]'''
	tags = ["key", "Github"]

[[rules]]
	description = "LinkedIn Client ID"
	regex = '''(?i)linkedin(.{0,20})?(?-i)['\"][0-9a-z]{12}['\"]'''
	tags = ["client", "LinkedIn"]

[[rules]]
	description = "LinkedIn Secret Key"
	regex = '''(?i)linkedin(.{0,20})?['\"][0-9a-z]{16}['\"]'''
	tags = ["secret", "LinkedIn"]

[[rules]]
	description = "Slack"
	regex = '''xox[baprs]-([0-9a-zA-Z]{10,48})?'''
	tags = ["key", "Slack"]

[[rules]]
	description = "EC"
	regex = '''-----BEGIN EC PRIVATE KEY-----'''
	tags = ["key", "EC"]


[[rules]]
	description = "Google API key"
	regex = '''AIza[0-9A-Za-z\\-_]{35}'''
	tags = ["key", "Google"]


[[rules]]
	description = "Heroku API key"
	regex = '''(?i)heroku(.{0,20})?['"][0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}['"]'''
	tags = ["key", "Heroku"]

[[rules]]
	description = "MailChimp API key"
	regex = '''(?i)(mailchimp|mc)(.{0,20})?['"][0-9a-f]{32}-us[0-9]{1,2}['"]'''
	tags = ["key", "Mailchimp"]

[[rules]]
	description = "Mailgun API key"
	regex = '''(?i)(mailgun|mg)(.{0,20})?['"][0-9a-z]{32}['"]'''
	tags = ["key", "Mailgun"]

[[rules]]
	description = "PayPal Braintree access token"
	regex = '''access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}'''
	tags = ["key", "Paypal"]

[[rules]]
	description = "Picatic API key"
	regex = '''sk_live_[0-9a-z]{32}'''
	tags = ["key", "Picatic"]

[[rules]]
	description = "Slack Webhook"
	regex = '''https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}'''
	tags = ["key", "slack"]

[[rules]]
	description = "Stripe API key"
	regex = '''(?i)stripe(.{0,20})?['\"][sk|rk]_live_[0-9a-zA-Z]{24}'''
	tags = ["key", "Stripe"]

[[rules]]
	description = "Square access token"
	regex = '''sq0atp-[0-9A-Za-z\-_]{22}'''
	tags = ["key", "square"]

[[rules]]
	description = "Square OAuth secret"
	regex = '''sq0csp-[0-9A-Za-z\\-_]{43}'''
	tags = ["key", "square"]

[[rules]]
	description = "Twilio API key"
	regex = '''(?i)twilio(.{0,20})?['\"][0-9a-f]{32}['\"]'''
	tags = ["key", "twilio"]

[[rules]]
	description = "Env Var"
	regex = '''(?i)(apikey|secret|key|api|password|pass|pw|host)=[0-9a-zA-Z-_.{}]{4,120}'''

[[rules]]
	description = "Port"
	regex = '''(?i)port(.{0,4})?[0-9]{1,10}'''
	[rules.allowlist]
		regexes = ['''(?i)port ''']
		description = "ignore export "



[[rules]]
	description = "Email"
	regex = '''[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}'''
	tags = ["email"]
	[rules.allowlist]
		files = ['''(?i)bashrc''']
		description = "ignore bashrc emails"


[[rules]]
	description = "Generic Credential"
	regex = '''(?i)(dbpasswd|dbuser|dbname|dbhost|api_key|apikey|secret|key|api|password|user|guid|hostname|pw|auth)(.{0,20})?['|"]([0-9a-zA-Z-_\/+!{}/=]{4,120})['|"]'''
	tags = ["key", "API", "generic"]
	# ignore leaks with specific identifiers like slack and aws
	[rules.allowlist]
		description = "ignore slack, mailchimp, aws"
		regexes = [
		    '''xox[baprs]-([0-9a-zA-Z]{10,48})''',
		    '''(?i)(.{0,20})?['"][0-9a-f]{32}-us[0-9]{1,2}['"]''',
		    '''(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}'''
		]

[[rules]]
	description = "High Entropy"
	regex = '''[0-9a-zA-Z-_!{}/=]{4,120}'''
  	file = '''(?i)(dump.sql|high-entropy-misc.txt)$'''
	tags = ["entropy"]
    [[rules.Entropies]]
        Min = "4.3"
        Max = "7.0"
    [rules.allowlist]
        description = "ignore ssh key and pems"
        files = ['''(pem|ppk|env)$''']
        paths = ['''(.*)?ssh''']

[[rules]]
	description = "Potential bash var"
	regex='''(?i)(=)([0-9a-zA-Z-_!{}=]{4,120})'''
	tags = ["key", "bash", "API", "generic"]
        [[rules.Entropies]]
            Min = "3.5"
            Max = "4.5"
            Group = "1"

[[rules]]
	description = "WP-Config"
	regex='''define(.{0,20})?(DB_CHARSET|NONCE_SALT|LOGGED_IN_SALT|AUTH_SALT|NONCE_KEY|DB_HOST|DB_PASSWORD|AUTH_KEY|SECURE_AUTH_KEY|LOGGED_IN_KEY|DB_NAME|DB_USER)(.{0,20})?['|"].{10,120}['|"]'''
	tags = ["key", "API", "generic"]

[[rules]]
	description = "Files with keys and credentials"
    file = '''(?i)(id_rsa|passwd|id_rsa.pub|pgpass|pem|key|shadow)'''

# Global allowlist
[allowlist]
	description = "image allowlists"
	files = ['''(.*?)(jpg|gif|doc|pdf|bin)$''']


`
