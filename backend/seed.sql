-- INSERT TEAMS --
INSERT INTO team (id, name, display_name, description) VALUES
(1, 'event-management', 'Event Management', 'Handles event creation and scheduling'),
(2, 'booking', 'Booking', 'Handles booking of artists and venues'),
(3, 'public-relations', 'Public Relations', 'Manages public-facing content'),
(4, 'admin', 'Admin', 'Full access to all resources'),
(5, 'member', 'Member', 'Basic view access');


-- INSERT PERMISSIONS --
INSERT INTO permission (id, name, display_name, description) VALUES
(1, 'view:event', 'View Event', 'Allows user to view event'),
(2, 'edit:event', 'Edit Event', 'Allows user to edit event'),
(3, 'delete:event', 'Delete Event', 'Allows user to delete event'),

(4, 'view:concert', 'View Concert', 'Allows user to view concert'),
(5, 'edit:concert', 'Edit Concert', 'Allows user to edit concert'),
(6, 'delete:concert', 'Delete Concert', 'Allows user to delete concert'),

(7, 'view:venue', 'View Venue', 'Allows user to view venue'),
(8, 'edit:venue', 'Edit Venue', 'Allows user to edit venue'),
(9, 'delete:venue', 'Delete Venue', 'Allows user to delete venue'),

(10, 'view:artist', 'View Artist', 'Allows user to view artist'),
(11, 'edit:artist', 'Edit Artist', 'Allows user to edit artist'),
(12, 'delete:artist', 'Delete Artist', 'Allows user to delete artist'),

(13, 'view:member', 'View Member', 'Allows user to view member'),
(14, 'edit:member', 'Edit Member', 'Allows user to edit member'),
(15, 'delete:member', 'Delete Member', 'Allows user to delete member'),

(16, 'view:team', 'View Team', 'Allows user to view team'),
(17, 'edit:team', 'Edit Team', 'Allows user to edit team'),
(18, 'delete:team', 'Delete Team', 'Allows user to delete team');


-- ASSIGN PERMISSIONS TO TEAMS --
INSERT INTO teams_permissions (team_id, permission_id)
SELECT 4, id FROM permission;

-- Member (team_id 5): only view permissions
INSERT INTO teams_permissions (team_id, permission_id)
SELECT 5, id FROM permission WHERE name LIKE 'view:%';

-- Event Management (team_id 1): view/edit event & concert
INSERT INTO teams_permissions (team_id, permission_id)
SELECT 1, id FROM permission WHERE name IN (
  'view:event', 'edit:event', 'view:concert', 'edit:concert'
);

-- Booking (team_id 2): view/edit venue & artist
INSERT INTO teams_permissions (team_id, permission_id)
SELECT 2, id FROM permission WHERE name IN (
  'view:venue', 'edit:venue', 'view:artist', 'edit:artist'
);

-- Public Relations (team_id 3): view event & artist
INSERT INTO teams_permissions (team_id, permission_id)
SELECT 3, id FROM permission WHERE name IN (
  'view:event', 'view:artist'
);
