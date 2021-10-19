#helm install armpgsql groundhog2k/postgres

helm install pgsql groundhog2k/postgres \
--set userDatabase.name=lemi011b \
--set userDatabase.user=postgres \
--set userDatabase.password=password \
