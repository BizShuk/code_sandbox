# Usage:
# ip, status_code = parse_line(line)
def parse_line(line):
    parts = line.split()
    return (parts[0], int(parts[8]))


class AccessLog(object):
    def __init__(self, path_to_logfile):
        # TODO: Implement

        self.ip_set = set([])
        self.status_count = {}
        # open file , do parse_line and save to data structure 
        # easier to  check which ip i have and count status.
        # data structure set for ip and dictionary for status.
        f = open('access.log', 'r')
        while 1:
            row = f.readline()
            if not row:
                break
            ip, status = parse_line(row)
            self.ip_set.add(ip)
            if not status in self.status_count.keys():
                self.status_count[status] = 0
            self.status_count[status] += 1

#        print self.ip_set
#        print self.status_count

#raise NotImplementedError()

    def contains_ip(self, ip):
        """Indicate whether the specified IP address
        is listed in the log file."""
        # TODO: Implement
        return ip in self.ip_set

    #raise NotImplementedError()

    def get_status_count(self, status_code):
        """Return the number of occurrances of the
        specified HTTP status code in the log file.
        """
        # TODO: Implement

        return self.status_count[
            status_code] if status_code in self.status_count.keys() else 0
    # raise NotImplementedError()

    #
    #ins = AccessLog('access.log')
    #
    #ip = ins.contains_ip('14.1.79.172')
    #ip2 = ins.contains_ip('0.0.0.0')
    #ip3 = ins.contains_ip('10.100.100.10')
    #status_count = ins.get_status_count(200)

    #print ip
    #print ip2
    #print ip3
    #print status_count
