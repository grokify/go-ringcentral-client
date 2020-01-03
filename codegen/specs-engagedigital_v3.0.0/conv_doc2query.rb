#!ruby

require 'multi_json'

doc = 'name: Team name.
leader_ids[]: List of user id as leaders | array:string
user_ids[]: List of user id as team members. | array:string'

doc = 'label: Folder’s label (mandatory).
parent_id: ID of the parent folder.
position: position of the folder. | integer
query: query of the folder as described in ​Search API documentation.​\n\nExample: “​active_and_assigned_to_me:true”
render_threads_count: boolean describing display of the number of threads. | boolean
role_restriction[only][]: list of roles allowed to see this folder. This parameter has to be a hash otherwise it will raise a 400 error. The key should be "only". For example: `&role_restriction[only][]=4e5596cdae70f677b5000002` | array:string
team_restriction[only][]: list of teams allowed to see this folder. Same thing as role_restriction: team_restriction parameter has to be a hash with the key "only". | array:string'

doc = '● name: Source name
● active: Activate/deactivate the source (Boolean)
● channel_id: Channel
● color: Color of the icon (see S​ ource colors​) (Integer)
● sla_response: Response time (seconds)
● sla_expired_strategy: SLA expired strategy ("max", "half" or "base")
● intervention_messages_boost: Priority boost of messages with intervention (Integer)
● transferred_tasks_boost: Priority boost of transferred tasks (Integer)
● hidden_from_stats: Hide from statistics (Boolean)
● default_category_ids[]: Default categories | array:string
● user_thread_default_category_ids[]: Default categories (agent messages) | array:string
● default_content_language: Default content language
● auto_detect_content_language: Auto-detect content language (Boolean)
● content_archiving: Automatic archiving of old contents (Boolean)
● content_archiving_period: Archive contents older than (seconds)'

doc = '● community_id: To filter identities on given community id.
● identity_group_id: To filter on given group id.
● user_id: To filter identities on given user id.
● sort: To change the criteria chosen to sort the identities. The value can be “created_at” or
“updated_at”.
● foreign_id: To filter identities on given user id
● uuid: To filter identities on given uuid'

doc = '● category_ids: User list of category ids (multiple).
● email: User email (mandatory).
● enabled: Whether the user is enabled or not (boolean).
● external_id: User external id.
● firstname: User firstname (mandatory).
● gender: User gender ("man" or "woman").
● identity_ids: User list of identity ids (multiple).
● lastname: User lastname (mandatory).
● locale: Language for the user interface.
● nickname: User nickname.
● role_id: User role id (mandatory).
● team_ids: User list of team ids (multiple).
● timezone: Use the timezone endpoint to get the timezone name (String), default is empty for
domain timezone.
● spoken_languages: List of locales corresponding to the languages spoken by the user (multiple).'

doc = '
● access_help_center
● access_previous_messages
● access_pull_mode
● admin_stamp_answer
● approve_content
● assign_intervention
● author_block_content
● close_content_thread
● create_and_destroy_extension
● create_community
● create_content_source
● create_user
● delay_export_content
● delete_content_thread
● impersonate_user
● invite_user
● manage_api_access_tokens
● manage_app_sdk_applications
● manage_automatic_exports_tasks *
● manage_categories
● manage_chat
● manage_custom_fields
● manage_custom_notifications
● manage_emails_templates
● manage_folders
● manage_ice
● manage_identities
● manage_own_notifications
● manage_reply_assistant *
● manage_roles
● manage_rules_engine_rules *
● manage_surveys *
● manage_tags
● manage_teams
● manage_topologies
● manage_users_of_my_teams
● monitor_tasks
● monitor_team_tasks
● mute_content
● open_content_thread
● publish_content
● read_community
● read_content_source
● read_event
● read_export
● read_identity
● read_own_stats
● read_presence
● read_stats
● read_surveys *
● read_user
● receive_tasks
● reply_with_assistant *
● search_contents
● search_event
● update_community
● update_content_source
● update_extension
● update_identity
● update_intervention
● update_own_intervention
● update_settings
● update_time_sheet
● update_user
● use_emoji
'

doc = '
● category_ids[]: User list of category ids (multiple).
● email: User email (mandatory).
● enabled: Whether the user is enabled or not (boolean).
● external_id: User external id, used for SSO.
● firstname: User firstname (mandatory).
● gender: User gender ("man" or "woman").
● identity_ids[]: User list of identity ids (multiple).
● lastname: User lastname (mandatory).
● locale: Language for the user interface.
● nickname: User nickname.
● password: User plain password (mandatory).
● role_id: User role id (mandatory).
● team_ids[]: User list of team ids (multiple).
● timezone: Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone.
● spoken_languages[]: List of locales corresponding to the languages spoken by the user (multiple).
'

doc = '
● name: Category name.
● parent_id: ID of parent category.
● color: displayed color for the category, see Category colors. | integer
● mandatory: mandatory categorization (Boolean).
● multiple: allow to assign multiple child categories (Boolean).
● post_qualification: post qualification (Boolean).
● unselectable: root category is unselectable (Boolean).
● source_ids[]: List of source id. | array:string
'

doc = '
● name: The name of the channel.
● source_ids: An array containing id of each source assigned to a channel (multiple). | array:string
● soft_capability: Number of tasks that can be assigned to agent by the routing before they are considered "occupied". | integer
● hard_capability: M​aximum number of tasks that can be assigned to agents. | integer
● task_timeout_seconds: this field defines the time before a task expires (in seconds). | integer
'

doc = '
● associated_type_name: The associated type of custom field. It can be IdentityGroup or Intervention.
● label: The label of the custom field (mandatory).
● key: The key of the custom field (example: customer_id). This is used to determine how it is
stored on identity groups.
● type: The type of the custom field. It can be string, boolean, text, integer, float, single_choice,
or multiple_choice (default: string).
● choices: A list of choices to be for single_choice, or multiple_choice types. This must be given
as array. | array:string
● multiple: true or false, this as no effect on single_choice, multiple_choice or boolean types
(default: false). | boolean
● position: an integer that indicates custom field’s position between others (default: -1). | integer
'

doc = '● company: Identity company.
● custom_field_values[custom_field_key]: Identity custom field with key « custom_field_key ». It
can be multiple if custom field is multiple or is has multiple_choice type.
● emails: Identity emails (multiple).
● firstname: Identity firstname.
● gender: Identity’s gender. It can be "man", "woman" or empty.
● home_phones: Identity home phones (mutiple).
● lastname: Identity lastname.
● mobile_phones: Identity mobile phones (multiple).
● notes: Identity notes.
● tag_ids: Identity tag ids (multiple).'

doc = '● agent_ids: List of agents to transfer the task to (multiple).
● bypass: Force the transfer to the first agent in agent_ids if set. When bypass is used,
agent_ids need to be provided and the other parameters will be ignored.
● category_ids: Filter agents receiving the task depending on their categories.
● language: Filter agents receiving the task depending on their spoken languages.
● team_ids: Filter agents receiving the task depending on their teams.
● comment: Add a comment to the task.'

doc = '● active: true or false, this field is used to enable/disable a webhook.
● label: The label of the webhook (mandatory).
● staging_use: true or false, this field is used to determine if a webhook will be run in staging
or production.
● url: The url of a webhook. This is used to determine the endpoint of your webhook, where
the data will be submitted.
● verify_token: The token used in your webhook.
● secret: The secret key that will be served as a ​X-Dimelo-Secret​ header in every request.
● registered_events: An array containing all the events that your webhook wants to subscribe.
'

doc = '● activity_presence_threshold: (in hours). | integer
● activity_tracking: ​Enable activity tracking (Boolean)
● beginning_of_week: (Day of week)
● category_tagging: A​ ctivate the forced categorization by source.​ (Boolean)
● content_languages: (See format)
● dashboard: Activate the dashboard (Boolean)
● deny_iframe_integration: Prevent the DD to be embed by other websites (Boolean)
● disable_password_autocomplete: (Boolean)
● expire_password_after: password expiration delay (in seconds) | integer
● expire_password_enabled: enable password expiration (Boolean)
● export_in_seconds: provide durations in seconds in export (Boolean)
● fold_useless_contents: fold archived contents (Boolean)
● fte_duration: FTE data period (in hours) | integer
● identity_merge: enable identity merge (Boolean)
● intervention_defer_rates[]: (Array of times in seconds) | array:integer
● intervention_defer_threshold: (in seconds) | integer
● intervention_rates: (Array of times in seconds) | array:integer
● locale: locale code (String)
● multi_lang: activate multi language support for messages (Boolean)
● name: Name of the Dimelo Digital (String)
● password_archivable_enabled: prohibit reusing old passwords (Boolean)
● password_archivable_size: number of archived passwords | integer
● password_min_length: minimum character length | integer
● password_non_word: should contain at least 1 non alphanumeric char (Boolean)
● password_numbers: should contain at least 1 number (Boolean)
● password_recovery_disabled: disable password recovery by email (Boolean)
● push_enabled: Enable push mode (Boolean)
● reply_as_any_identity: Enable reply as any identity (Boolean)
● rtl_support: Enable right to left support (Boolean)
● self_approval_required: ​Allow authors to ask approval of their messages (Boolean)
● session_timeout: Session timeout (in minutes) | integer
● spellchecking: Enable spellchecking (Boolean)
● style: Defines the DD’s design (String)
● third_party_services_disabled: Disable third-party services (tracking...) (Boolean)
● timezone: Use the timezone endpoint to get the timezone name (String)
● track_js: Track JS errors (Boolean)
● type: Can be ‘demo’, ‘production’ or ‘archived’
● urgent_task_threshold: Chat max response time (in seconds) | integer
● use_system_font: Experimental (Boolean)'

doc = 'active: true or false, this field is used to enable/disable a time sheet. | boolean
● label: The label of the time sheet (mandatory).
● source_ids: An array containing id of each source using your time sheet. | array:string
● holidays_region: A string containing the first two letters of your country (example: "fr"/"en"/"es"), useful to bootstrap default holidays following to a country.
● holidays: An array containing one or more hash of holidays, a holiday must contain a name (string) and a date (string), the date must be in a valid format, a valid format is a format corresponding to your domain’s locale).
● monday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. For example: “a-b,c-d”: “a” is the beginning of the first interval of the day, “b” is the ending of the first interval of the day, “c” is the beginning of the second interval of the day, “d” is the ending of the second interval of the day
● tuesday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
● wednesday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
● thursday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
● friday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
● saturday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
● sunday_hours: this field define the time intervals of the day (in secs). An empty string means that there are no business hours on this day. See `monday_hours` for the format.
'
parameters = []

doc.split("\n").each_with_index do |line,i|
  puts "LINE[#{i}] " + line
  if line =~ /^●?\s*([\w\[\]]+)\s*:\s*(.+)\s*$/
    name = $1
    desc = $2.strip
    type = ''
    format = ''
    items = {}
    param = {name: name, in: 'query', required: false}
    if desc =~ /^\s*(.+)\s*\|\s*(.+):(.+)\s*$/
      desc = $1.strip
      type = $2
      if type == 'array'
        items = {type: $3}
      end
    elsif desc =~ /^\s*(.+)\s*\|\s*(.+)\s*\.?\s*$/
      desc = $1
      type = $2
    elsif desc =~/\((\w+)\)/
      type_raw = $1.strip.downcase
      if type_raw == 'seconds'
        type = 'integer'
      elsif type_raw == 'mandatory'
        type = 'string'
        param[:required] = true
      elsif type_raw == 'multiple'
        type = 'array'
        items = {type: 'string'}
      else 
        type = type_raw
      end
    end
    param[:required] = false
    desc.gsub!(/\s*\(mandatory\)/,'')
    param[:description] = desc.strip
    if type.length > 0
      param[:schema] = {} unless param.key? :schema
      param[:schema][:type] = type
    else
        param[:schema] = {} unless param.key? :schema
        param[:schema][:type] = 'string'
    end
    if format.length > 0
      param[:format] = format
    end
    if type == 'array'
      if items.length >0 
        param[:schema][:items] = items
      end
    end
    parameters.push param
  elsif line =~ /^●?\s*(\w+)(\s*\*\s*)?$/
    name = $1.strip
    desc = $2 || ''
    desc.strip!
    param = {name: name, in: 'query', required: false}
    if desc == "*"
      param[:description] = "permission only available with the corresponding extension enabled"
    end
    param[:schema] = {} unless param.key? :schema
    param[:schema][:type] = 'boolean'
    parameters.push param
  else
    puts "E_NO_MATCH: " + line
  end
end

puts doc

puts MultiJson.encode parameters, pretty: true