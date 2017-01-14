#!/usr/bin/env bash

NAME=domio_public
DESC="Domio Public service"
PIDFILE="/var/run/${NAME}.pid"
LOGFILE="/var/log/${NAME}.log"

DAEMON="/user/local/bin/domio_public"

#start-stop-daemon --start --background --make-pidfile --pidfile /var/run/domio.pid --exec /domio/domio
#start-stop-daemon --stop --pidfile /var/run/domio.pid

START_OPTS="--start --background --make-pidfile --pidfile ${PIDFILE} --exec ${DAEMON} start"
STOP_OPTS="--stop --pidfile ${PIDFILE}"

test -x $DAEMON || exit 0

set -e


case "$1" in
 start)
# echo -----------------------------------
# echo "PIDFILE: ${PIDFILE}"
# echo "NAME: ${NAME}"
# echo -n "Starting ${DESC}: "
# echo -----------------------------------
# start-stop-daemon $START_OPTS >> $LOGFILE
 echo "$NAME."
 start-stop-daemon --start --background --make-pidfile --pidfile /var/run/domio_public.pid --exec /usr/local/bin/domio_public start
 ;;
stop)
# echo -----------------------------------
# echo "PIDFILE: ${PIDFILE}"
# echo "NAME: ${NAME}"
# echo -n "Stopping $DESC: "
# echo -----------------------------------
# start-stop-daemon $STOP_OPTS
# echo "$NAME."
# rm -f $PIDFILE
 start-stop-daemon --stop --pidfile /var/run/domio_public.pid
 ;;
restart|force-reload)
 echo -n "Restarting $DESC: "
 start-stop-daemon $STOP_OPTS
 sleep 1
 start-stop-daemon $START_OPTS >> $LOGFILE
 echo "$NAME."
 ;;
*)
 N=/etc/init.d/$NAME
 echo "Usage: $N {start|stop|restart|force-reload}" >&2
 exit 1
 ;;
esac

exit 0